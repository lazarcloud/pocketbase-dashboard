<script>
  import "./global.css"
  import "./styles.css"
  import { Navbar } from "$lib/ui"
</script>

<Navbar />
<div class="main">
  <slot />
</div>

<style>
  .main {
    min-height: 100vh;
    padding: 1rem;
  }
</style>
