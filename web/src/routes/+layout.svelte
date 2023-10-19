<script>
  import "./styles.css"
  import "./base.css"
</script>

<div class="main">
  <slot />
</div>

<style>
  .main {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }
</style>
