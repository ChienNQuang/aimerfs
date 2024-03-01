export function Ls(dir: string) {
  return fetch(`http://localhost:8080/api/ls?dir=${dir}`)
    .then((res) => res.json())
    .then((data) => data as string[])
}
