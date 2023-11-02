<script>
  import "./styles.css"
  import "./base.css"
  import { auth } from "$lib/auth"
  let password
</script>

<div class="main">
  {#if $auth.password != null}
    <slot />
  {:else}
    <div class="block login">
      <h1>Log in</h1>
      {#if $auth.error}
        <p style="color: red;">{$auth.error}</p>
      {/if}
      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        on:keyup={(e) => {
          if (e.key == "Enter")
            auth.set({
              password,
              error: null,
            })
        }}
      />
      <button
        on:click={() =>
          auth.set({
            password,
            error: null,
          })}>Log in</button
      >
    </div>
  {/if}
</div>

<style>
  .main {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }
  h1 {
    margin-bottom: 1rem;
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

    text-align: center;
    width: 100%;
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
  input[type="password"] {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #d7dde4;
    border-radius: 4px;
    margin-bottom: 1rem;
  }
  .block {
    border: 1px solid #d7dde4;
    background-color: var(--baseColor);
    padding: 1rem;
  }
  .login {
    width: clamp(300px, 50%, 500px);
    margin: auto;
  }
</style>
