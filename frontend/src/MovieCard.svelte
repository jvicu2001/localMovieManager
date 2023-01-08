<script lang="ts">
    export let poster_path: string = ""
    export let adult: boolean = false
    export let overview: string = ""
    export let release_date: string = ""
    export let genre_ids: number[] = []
    export let id: number
    export let original_title: string = ""
    export let original_language: string = ""
    export let title: string = ""
    export let backdrop_path: string = ""
    export let popularity: number = 0
    export let vote_count: number = 0
    export let video: boolean = false
    export let vote_average: number = 0


</script>

<div class="relative bg-cover overflow-hidden rounded-lg max-w-sm transform hover:scale-110 ease-in-out duration-75 hover:z-20 hover:shadow-md">
    <img src={
        poster_path ? "https://image.tmdb.org/t/p/w500" + poster_path : "https://via.placeholder.com/500x750"
    } alt="" class="inset-0 w-full h-full object-cover" />

    <div class="absolute inset-0 align-bottom">
        <h3 class="text-white text-2xl font-bold p-4 bg-gradient-to-b from-black">{title}</h3>
    </div>
</div>

<style>
    /* your styles go here */
</style>