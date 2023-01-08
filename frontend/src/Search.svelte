<script lang="ts">
    import {SearchMovies} from '../wailsjs/go/modules/SearchStruct.js'
    import MovieCard from './MovieCard.svelte';
    import type {modules} from "../wailsjs/go/models";
    
    // Icons
    import Icon from '@iconify/svelte';
    
    import Pagination from './Pagination.svelte';

    async function searchMovies(
        query: string,
        page: number = 1,
        include_adult: boolean = false,
        region: string = "US",
        language: string = "en-US"
    ) {
        var options: modules.SearchMoviesOptions = {
            page: page,
            include_adult: include_adult,
            region: region,
            language: language,
        }
        let unmarshalledJson: string = await SearchMovies(query,options)
        queryResult = JSON.parse(unmarshalledJson)
        queryMovies = queryResult["Results"]
    }

    let searchText: string
    let queryResult: JSON
    let queryMovies: JSON[] = []
</script>

<div class="container mx-auto">
    <div class="row py-3">
        <div class="col-12">
            <input type="text text-red-200" bind:value={searchText} />
            <button on:click={() => searchMovies(searchText)} class="text-white ring-2 rounded-md p-1"><Icon icon="material-symbols:search"/></button>
        </div>
    </div>

    <!-- Pagination -->
    {#if queryResult}
    <div class="row py-3">
        <div class="col-12">
            <Pagination
                currentPage = {queryResult["Page"]}
                totalPages = {queryResult["total_pages"]}
                searchText = {searchText}
                searchMovies = {searchMovies}
            />
        </div>
    </div>
    {/if}

    <div class="row">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-2 px-4">
            {#each queryMovies as movie}
                <MovieCard
                    adult = {movie["adult"]}
                    backdrop_path = {movie["backdrop_path"]}
                    id = {movie["id"]}
                    original_title = {movie["original_title"]}
                    genre_ids = {movie["genre_ids"]}
                    popularity = {movie["popularity"]}
                    poster_path = {movie["poster_path"]}
                    release_date = {movie["release_date"]}
                    title = {movie["title"]}
                    overview = {movie["overview"]}
                    video = {movie["video"]}
                    vote_average = {movie["vote_average"]}
                    vote_count = {movie["vote_count"]}
                />
            {/each}
        </div>
    </div>

    <!-- Pagination -->
    {#if queryResult}
    <div class="row py-3">
        <div class="col-12">
            <Pagination
                currentPage = {queryResult["Page"]}
                totalPages = {queryResult["total_pages"]}
                searchText = {searchText}
                searchMovies = {searchMovies}
            />
        </div>
    </div>
    {/if}
</div>