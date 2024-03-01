<script setup lang="ts">
import { computed, ref, type Ref } from 'vue'
import { Ls } from '../api/index'

const { currentDirectory, oldCommands, handleCommand } = useFileSystemShell()
const { input, clearInput, suggestions } = useFileSystemInput()
const showSuggestions = ref(false)

function useFileSystemShell() {
  const currentDirectory = ref('/')
  const oldCommands = ref<string[]>([])

  const handleSubmit = (text: string) => {
    oldCommands.value.push(`${currentDirectory.value} > ${text}`)
  }

  const handleCommand = async (command: string) => {
    const [commandName, ...args] = command.split(' ')

    switch (commandName) {
      case 'cd':
        if (args[0] === '..') {
          currentDirectory.value = currentDirectory.value.split('/').slice(0, -1).join('/')
        } else {
          currentDirectory.value = `${currentDirectory.value}/${args[0]}`
        }
        break
      case 'ls':
        if (args[0] !== undefined) {
          const items = await Ls(args[0])
          oldCommands.value.push(items.join(' '))
        } else {
          const items = await Ls(currentDirectory.value)
          oldCommands.value.push(items.join(' '))
        }
        break
      case 'mkdir':
        break
      case 'touch':
        break
      default:
        break
    }
  }

  return {
    currentDirectory,
    oldCommands,
    handleSubmit,
    handleCommand
  }
}

function useFileSystemInput() {
  const input = ref('')
  const commands = ['cd', 'ls', 'mkdir', 'touch']

  const isTypingCommand = computed(() => {
    return !input.value.includes(' ')
  })
  const suggestions = computed(() => {
    if (input.value === '') {
      return []
    }

    if (isTypingCommand.value) {
      return commands.filter((command) => command.startsWith(input.value))
    }

    return []
  })

  function clearInput() {
    input.value = ''
  }

  return {
    input,
    suggestions,
    clearInput
  }
}
</script>

<template>
  <div class="border h-full overflow-auto">
    <div v-for="command in oldCommands" :key="command">
      {{ command }}
    </div>
    <div class="flex items-center">
      <span>{{ currentDirectory }} ></span>
      <div class="relative">
        <input
          v-model="input"
          @keyup.enter="
            () => {
              handleCommand(input)
              clearInput()
            }
          "
          @focusin="showSuggestions = true"
          @focusout="showSuggestions = false"
        />
        <div class="absolute border px-2 py-1" v-if="suggestions.length > 0 && showSuggestions">
          <div v-for="suggestion in suggestions" :key="suggestion">
            <div @click="input = suggestion" class="hover:cursor-pointer">
              {{ suggestion }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
