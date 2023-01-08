<script lang="ts">
    export let profile_path : string = ""
    export let adult : boolean = false
    export let id : number
    export let name : string = ""
    export let popularity : number = 0
    export let known_for : any[] = []


</script>

<div class="relative bg-cover overflow-hidden rounded-lg max-w-sm transform hover:scale-110 ease-in-out duration-75 hover:z-20 hover:shadow-md">
    <img src={
        profile_path ? "https://image.tmdb.org/t/p/w500" + profile_path : "https://via.placeholder.com/500x750"
    } alt="" class="inset-0 w-full h-full object-cover" />

    <div class="absolute inset-0 align-bottom">
        <h3 class="text-white text-2xl font-bold p-4 bg-gradient-to-b from-black">{name}</h3>
    </div>
</div>

<style>
    /* your styles go here */
</style>