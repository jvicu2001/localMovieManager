<script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {Greet, TestAPI, UpdateAPIkey} from '../wailsjs/go/main/App.js'

  let resultText: string = "Please enter your name below 👇"
  let name: string

  let APIKey: string
  let queryResult: string = "Placeholder"

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }

  function updateAPIKey(key: string) {
    UpdateAPIkey(key)
    testAPI("Bee Movie")
  }

  function testAPI(query: string) {
    TestAPI(query).then(result => queryResult = result)
  }
  testAPI("Bee Movie")
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input shadow-sm text-black" id="name" type="text"/>
    <button class="btn" on:click={greet}>Greet</button>
  </div>
  <div class="py-5 px-5" id="api_input">
    <input autocomplete="off" bind:value={APIKey} class="input shadow-sm text-black" id="api_key" type="text"/>
    <button class="btn rounded-lg hover:bg-sky-900 p-1" on:click={testAPI(APIKey)}>Change API Key</button>
  </div>
  <h3>{queryResult}</h3>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>

