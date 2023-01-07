<script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {Greet, TestAPI} from '../wailsjs/go/main/App.js'
  import Search from './Search.svelte'

  let resultText: string = "Please enter your name below 👇"
  let name: string

  let queryResult: string = "Placeholder"

  function greet(): void {
    Greet(name).then(result => resultText = result)
  }

  function testAPI(query: string) {
    TestAPI(query).then(result => queryResult = result)
  }
  testAPI("Bee Movie")
</script>

<main class="">
  <Search/>
  <!--<h3>{queryResult}</h3>-->
</main>

<style>

</style>

