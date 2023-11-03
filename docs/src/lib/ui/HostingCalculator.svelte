<script>
  const imageName = "monsieurlazar/pocketbase-dashboard"
  let method = "dockerCompose"
  let containerName = "pocketbase-dashboard"
  let serverOrigin = "https://pocket.example.com"
  let password = "example"
  let storagePath = "/home/pocketbase/metadata"
</script>

<section>
  <div class="controls">
    <h2>Choose your hosting method</h2>

    <div class="buttons">
      <input
        type="radio"
        id="dockerRun"
        bind:group={method}
        value="dockerRun"
        name="method"
      />
      <label class="button" for="dockerRun">Docker Run Command</label>
      <input
        type="radio"
        id="dockerCompose"
        bind:group={method}
        value="dockerCompose"
        name="method"
      />
      <label class="button" for="dockerCompose">Docker Compose File</label>
    </div>

    <form>
      <label for="containerName">Docker Container Name</label>
      <input id="containerName" type="text" bind:value={containerName} />

      <label for="containerName">Server Origin Domain</label>
      <input id="containerName" type="text" bind:value={serverOrigin} />

      <label for="containerName">Password</label>
      <input id="containerName" type="text" bind:value={password} />

      <label for="containerName">Persistent Storage Path</label>
      <input id="containerName" type="text" bind:value={storagePath} />
    </form>
  </div>

  <div class="commands">
    <h3>Create docker network.</h3>
    <div>docker network create lazar-static</div>
    {#if method == "dockerRun"}
      <h3>Run with a docker run command.</h3>
      <p class="code">
        docker run -d -p 8081:80 -e ORIGIN={serverOrigin} -e DEFAULT_PASSWORD={password}
        --name
        {containerName} -v /var/run/docker.sock:/var/run/docker.sock -v {storagePath}:/data
        --network=lazar-static {imageName}
      </p>
    {:else if method == "dockerCompose"}
      <h3>Run with a docker compose file.</h3>
      <div class="code">
        <span style="--space: 0;">
          <span class="green">version</span>: <span class="blue">"3.8"</span>
        </span><br />
        <span style="--space: 0;">
          <span class="green">services</span>:
        </span><br />
        <span style="--space: 1;">
          <span class="green">lazar-dash</span>:
        </span><br />
        <span style="--space: 1;">
          <span class="green">image</span>:
          <span class="blue">{imageName}</span>
        </span><br />
        <span style="--space: 1;">
          <span class="green">container_name</span>:
          <span class="blue">{containerName}</span>
        </span><br />
        <span style="--space: 1;">
          <span class="green">environment</span>:
        </span><br />
        <span style="--space: 2;">
          - <span class="blue">ORIGIN={serverOrigin}</span>
        </span><br />
        <span style="--space: 2;">
          - <span class="blue">DEFAULT_PASSWORD={password}</span>
          <span class="gray">//defaults to password</span>
        </span><br />
        <span style="--space: 1;">
          <span class="green">volumes</span>:
        </span><br />
        <span style="--space: 2;">
          - <span class="blue">/var/run/docker.sock:/var/run/docker.sock</span>
        </span><br />
        <span style="--space: 2;">
          - <span class="blue">{storagePath}:/data</span>
        </span><br />
        <span style="--space: 1;">
          <span class="green">networks</span>:
        </span><br />
        <span style="--space: 2;">
          - <span class="blue">lazar-static</span>
        </span><br />
        <span style="--space: 1;">
          <span class="green">restart</span>: <span class="blue">always</span>
        </span><br />
        <span /><br />
        <span style="--space: 0;">
          <span class="green">networks</span>:
        </span><br />
        <span style="--space: 1;">
          <span class="green">lazar-static</span>:
        </span><br />
        <span style="--space: 2;">
          <span class="green">external</span>: <span class="blue">true</span>
        </span><br />
      </div>
    {/if}
  </div>
</section>

<style>
  span {
    --space: 0;
    margin-left: calc(16px * var(--space));
  }
  span.green {
    color: rgb(125, 240, 125);
  }
  span.blue {
    color: rgb(0, 247, 255);
  }
  span.gray {
    color: gray;
  }

  /* layouts */

  section {
    width: 100%;
    max-width: 1000px;
    margin: 0 auto;
    border: 1px solid #d7dde4;
    display: grid;
    grid-template-columns: 1fr 1fr;
    border-radius: 0.5rem;
    background-color: white;
    min-height: 100vh;
  }
  .controls,
  .commands {
    padding: 1rem 2rem;
  }
  .commands {
    border-left: 1px solid #d7dde4;
  }

  /* type selector */

  input[type="radio"] {
    display: none;
  }
  input[type="radio"]:checked + label {
    background-color: #d7dde4;
  }
  label.button {
    display: block;
    padding: 0.5rem;
    margin-left: -0.5rem;
    border-radius: 4px;
    cursor: pointer;
  }
  .buttons {
    display: flex;
    gap: 0.5rem;
    margin-bottom: 0.5rem;
  }

  /* input form */

  form {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: repeat(4, 1fr);
    gap: 1rem;
  }
  form label {
    display: flex;
    margin-bottom: 0.5rem;
    place-items: center;
  }
  form input {
    padding: 0.5rem;
    border-radius: 4px;
    background-color: transparent;
    border: 1px solid #d7dde4;
  }

  .code {
    padding: 1rem;
    border-radius: 0.5rem;
    margin: 1rem 0;
    border: 1px solid #d7dde4;
    width: 100%;
    overflow-x: auto;

    overflow-wrap: break-word;
    word-wrap: break-word;
    -ms-word-break: break-word;
    word-break: break-word;
  }

  @media (max-width: 900px) {
    section {
      grid-template-columns: 1fr;
    }
    .commands {
      border-left: none;
      border-top: 1px solid #d7dde4;
    }
  }
</style>
