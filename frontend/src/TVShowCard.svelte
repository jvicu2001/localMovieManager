<script lang="ts">
    export let poster_path: string = ""
    export let popularity: number = 0
    export let id: number
    export let backdrop_path: string = ""
    export let vote_average: number = 0
    export let overview: string = ""
    export let first_air_date: string = ""
    export let origin_country: string[] = []
    export let genre_ids: number[] = []
    export let original_language: string = ""
    export let vote_count: number = 0
    export let name: string = ""
    export let original_name: string = ""


</script>

<div class="relative bg-cover overflow-hidden rounded-lg max-w-sm transform hover:scale-110 ease-in-out duration-75 hover:z-20 hover:shadow-md">
    <img src={
        poster_path ? "https://image.tmdb.org/t/p/w500" + poster_path : "https://via.placeholder.com/500x750"
    } alt="" class="inset-0 w-full h-full object-cover" />

    <div class="absolute inset-0 align-bottom">
        <h3 class="text-white text-2xl font-bold p-4 bg-gradient-to-b from-black">{name}</h3>
    </div>
</div>

<style>
    /* your styles go here */
</style>