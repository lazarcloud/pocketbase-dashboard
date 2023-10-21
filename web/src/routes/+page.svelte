<script>
  import { goto, invalidateAll } from "$app/navigation"
  import Logo from "$lib/Logo.svelte"
  import { onMount } from "svelte"

  let baseURL = null

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

  function randomName(length = 8) {
    let name = ""
    const possible =
      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    for (let i = 0; i < length; i++)
      name += possible.charAt(Math.floor(Math.random() * possible.length))
    return name
  }
  async function getProjectsClient() {
    const res = await getProjects()
    projectsData = res
  }
  let instanceName = randomName()
  let projectsData = null
  let showPopup = false

  let interval = null

  onMount(() => {
    baseURL = window.location.href
    interval = setInterval(async () => {
      let newData = await getProjects()
      if (JSON.stringify(newData) !== JSON.stringify(projectsData)) {
        projectsData = newData
      }
    }, 1000)

    return () => {
      clearInterval(interval)
    }
  })
</script>

<section class="block">
  <nav>
    <Logo />
    <h1>Pocketbase Dashboard</h1>
    <p>by lazar</p>
  </nav>

  <div class="controls">
    <button>Settings</button>
    <button on:click={() => (showPopup = true)} class="black"
      ><div class="plus" />
      New Project</button
    >
  </div>
</section>
<section class="content">
  {#if baseURL != null}
    {#if projectsData == null}
      <p>No projects found</p>
    {:else}
      {#each projectsData as project}
        <div class="project">
          <h2>{project.Name} - {project.Status}</h2>
          <p>Project description</p>
          <p>Created: 2021-01-01</p>
          <p>API URL: https://pocket.lazar.lol/project/lazar/api/</p>
          <p>Dashnoard URL: https://pocket.lazar.lol/project/lazar/_/</p>
          <div class="separator" />
          <div class="controls">
            <button
              on:click={() => {
                goto(`/project/${project.Name.split("-")[1]}/_/`)
              }}>Go to Dashboard</button
            >
            <button> Edit</button>
          </div>
        </div>
      {/each}
    {/if}
  {:else}
    <p>Loading...</p>
  {/if}
</section>
{#if showPopup}
  <section
    class="overlay"
    on:click={() => {
      showPopup = false
    }}
  />
  <section class="popup">
    <div>
      <h1>New Project</h1>

      <div class="separator" />
      <label for="instanceName">Name</label>
      <input
        type="text"
        id="instanceName"
        bind:value={instanceName}
        placeholder="Instance name"
      />
      <label>Description</label>
      <input type="text" placeholder="Instance description" />
    </div>
    <div>
      <div class="separator" />

      <div class="controls space">
        <p>Version: v.0.19</p>

        <div class="controls">
          <button on:click={() => (showPopup = false)}>Close</button>
          <button
            on:click={async () => {
              const res = await createProject(instanceName)
              console.log(res)
              showPopup = false
              projectsData = await getProjects()
              await invalidateAll()
            }}
            class="black"
          >
            Create</button
          >
        </div>
      </div>
    </div>
  </section>
{/if}

<style>
  nav {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 1rem;
  }
  button {
    background: #fff;
    border: 2px solid #000000;
    border-radius: 4px;
    padding: 8px 20px;
    font-size: 16px;
    cursor: pointer;
    font-weight: 500;
    transition: background-color 0.2s ease-in-out;
    display: flex;
  }
  button.black {
    background: #000000;
    color: #fff;
  }
  button:hover {
    background: #edf0f3;
  }
  button.black:hover {
    background: #2c2c30;
  }
  .plus::before {
    content: "+";
    margin-right: 0.5rem;
  }
  .project {
    border: 1px solid #d7dde4;
    background-color: var(--baseColor);
    padding: 1rem;
    margin-bottom: 1rem;
    border-radius: 0.25rem;
  }
  .block {
    border: 1px solid #d7dde4;
    background-color: var(--baseColor);
    padding: 1rem;
  }
  .content {
    padding: 1rem;
    overflow-y: auto;
  }
  .controls {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    justify-content: flex-end;
  }
  .controls.space {
    justify-content: space-between;
  }
  .overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 50%;
    padding: 1rem;
    height: 100%;
    z-index: 100;
    background: rgba(0, 0, 0, 0.3);
  }
  .popup {
    position: absolute;
    top: 0;
    right: 0;
    width: 50%;
    padding: 3rem 2rem;
    height: 100%;
    background: var(--baseColor);
    z-index: 100;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .separator {
    border-bottom: 1px solid #d7dde4;
    margin-top: 0.5rem;
    margin-bottom: 1rem;
  }
  input[type="text"] {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #d7dde4;
    border-radius: 4px;
    margin-bottom: 1rem;
  }
</style>
