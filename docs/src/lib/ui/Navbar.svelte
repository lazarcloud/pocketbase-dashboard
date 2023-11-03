<script>
  import { page } from "$app/stores"
  import { Logo } from "$ui"
  let path = $page.url.pathname
  $: path = $page.url.pathname
</script>

<nav>
  <div class="item">
    <Logo />
    <p>PB Dash</p>
  </div>
  <div class="item">
    <a
      class={path == "/pocketbase-dashboard/" ? "selected" : ""}
      href="/pocketbase-dashboard/">Home</a
    >
    <a
      class={path == "/pocketbase-dashboard/selfhost" ? "selected" : ""}
      href="/pocketbase-dashboard/selfhost">Self-hosting</a
    >
    <a
      class={path == "/pocketbase-dashboard/roadmap" ? "selected" : ""}
      href="/pocketbase-dashboard/roadmap">Roadmap</a
    >
  </div>
</nav>

<style>
  nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
  }
  .item,
  a {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    cursor: pointer;
  }
  a {
    color: inherit;
    text-decoration: none;
  }
  a.selected,
  a:hover {
    text-decoration: underline;
  }
</style>
