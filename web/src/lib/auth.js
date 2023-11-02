import { writable } from "svelte/store"

const maxTime = 3600 // 1 hour
function persist(key, startValue) {
  let value = startValue

  if (typeof window !== "undefined") {
    const storedValue = document.cookie
      .split("; ")
      .find((row) => row.startsWith(key))
    value = storedValue ? JSON.parse(storedValue.split("=")[1]) : startValue
  }

  const store = writable(value)

  store.subscribe(($value) => {
    if (typeof window !== "undefined") {
      document.cookie = `${key}=${JSON.stringify(
        $value
      )}; max-age=${maxTime}; secure; samesite=strict`
    }
  })

  return store
}

export const auth = persist("auth", {
  password: null,
  error: "",
})
