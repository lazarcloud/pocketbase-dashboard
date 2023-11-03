<script>
  import { HostingCalculator } from "$lib/ui"
</script>

<svelte:head>
  <title>PB Dash Self-hosting</title>
</svelte:head>

<h1 class="title">Self-hosting Guide</h1>

<div class="section">
  <div class="container">
    <h2>Getting Started</h2>

    <p>Follow the steps below to set up PocketBase Dashboard using Docker.</p>

    <h3>Prerequisites</h3>

    <p>
      Make sure you have Docker installed on your system. If not, you can
      download and install it from the <a
        href="https://www.docker.com/get-started">official Docker website</a
      >.
    </p>
  </div>
</div>

<HostingCalculator />
