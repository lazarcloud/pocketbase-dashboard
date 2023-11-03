<script>
  import { page } from "$app/stores"
  import { Logo, Github } from "$ui"
  let path = $page.url.pathname
  $: path = $page.url.pathname
</script>

<nav>
  <div class="item">
    <Logo />
    <p>PB Dash</p>
  </div>
  <div class="item">
    <a class={path == "/" ? "selected" : ""} href="/">Home</a>
    <a class={path == "/selfhost" ? "selected" : ""} href="/selfhost"
      >Self-hosting</a
    >
    <a class={path == "/roadmap" ? "selected" : ""} href="/roadmap">Roadmap</a>
    <a href="/github">
      <Github />
    </a>
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
    gap: 2.5rem;
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
  @media (max-width: 600px) {
    .item:first-child p {
      display: none;
    }
  }
  @media (max-width: 520px) {
    nav {
      flex-direction: column;
      align-items: flex-start;
      align-items: center;
    }
    .item {
      margin-bottom: 1rem;
      flex-direction: column;
      gap: 1rem;
    }
  }
</style>
