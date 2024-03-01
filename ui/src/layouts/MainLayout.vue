<script setup lang="ts">
import AppIcon from '@/components/AppIcon.vue'
</script>

<template>
  <div class="h-screen">
    <header class="border-b-2 py-2 h-[10%]">
      <AppIcon />
    </header>
    <main class="h-[90%]">
      <slot />
    </main>
  </div>
</template>
