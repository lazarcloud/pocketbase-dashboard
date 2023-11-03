<script>
</script>

<svelte:head>
  <title>PB Dash Home</title>
</svelte:head>

<h1 class="title">Pocketbase Dashboard</h1>
<div class="section">
  <div class="container">
    <p>
      PocketBase Dashboard is a self-hosted solution that allows you to manage
      and use PocketBase for personal use. With this, you can have full control
      over your data and applications in a convenient and user-friendly way.
    </p>
  </div>
</div>

<section>
  <div class="container">
    <ul>
      <li>
        <i />
        Self-Hosting
      </li>
      <li>
        <i />
        User-Friendly Interface
      </li>
      <li>
        <i />
        Unlimited Projects
      </li>
    </ul>
  </div>
</section>
<div class="center">
  <img
    src="https://github.com/lazarcloud/pocketbase-dashboard/blob/main/images/app.jpeg?raw=true"
    alt="App Usage"
  />
  <p>Example App Usage</p>
</div>

<div class="center">
  <img
    src="https://github.com/lazarcloud/pocketbase-dashboard/blob/main/images/containers.png?raw=true"
    alt="Docker Containers"
  />
  <p>Generated Containers</p>
</div>

<style>
  .center {
    text-align: center;
  }
  img {
    width: 100%;
    max-width: 800px;
    margin: 1rem auto;
    box-shadow: 0 0 0 1px black;

    border-radius: 0.5rem;
  }
  ul {
    list-style: none;
    display: flex;
    justify-content: space-between;
    gap: 1rem;
    background-color: white;
    padding: 1rem;
    border-radius: 69rem;
  }
  li {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-weight: 700;
  }
  i::before {
    content: "✓";
    color: green;
  }
</style>
