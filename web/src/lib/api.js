let baseURL = "http://localhost:8080/"

export async function fetchAPI(url) {
  const response = await fetch(baseURL + url)
  const data = await response.json()
  return data
}

export async function getProjects() {
  return await fetchAPI("projects")
}
export async function createProject(name) {
  let response = await fetchAPI("create?slug=" + name)
  return {
    id: response,
    name: name,
  }
}
