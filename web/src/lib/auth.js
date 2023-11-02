import { writable } from "svelte/store"

export const auth = writable({
  password: null,
  error: "",
})
