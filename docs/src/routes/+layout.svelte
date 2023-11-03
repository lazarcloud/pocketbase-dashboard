<script>
  import "./global.css"
  import "./styles.css"
  import { Navbar, BottomBar } from "$lib/ui"
</script>

<Navbar />
<div class="main">
  <div class="bg" />
  <slot />
</div>
<BottomBar />

<style>
  .bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 90%;
    z-index: -10;
  }
  .main {
    min-height: 100vh;
    padding: 1rem;
  }
</style>
