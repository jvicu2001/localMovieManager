<script lang="ts">
    import {SearchMovies} from '../wailsjs/go/modules/SearchStruct.js'
    import MovieCard from './MovieCard.svelte';
    import {modules} from "../wailsjs/go/models";
    
    // Icons
    import Icon from '@iconify/svelte';

    async function searchMovies(query: string) {
        var options: modules.SearchMoviesOptions = {
            page: 1,
            include_adult: false,
            region: "US",
            language: "en-US",
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

    <div class="row">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-2 px-4">
            {#each queryMovies as movie}
                <MovieCard {...movie}/>
            {/each}
        </div>
    </div>
</div>