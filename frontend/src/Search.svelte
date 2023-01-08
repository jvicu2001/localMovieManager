<script lang="ts">
    import {SearchMovies, SearchTV, SearchPeople} from '../wailsjs/go/modules/SearchStruct.js'
    import MovieCard from './MovieCard.svelte';
    import TVShowCard from './TVShowCard.svelte';
    import PeopleCard from './PeopleCard.svelte';
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
        queryValues = queryResult["Results"]
    }

    async function searchTV(
        query: string,
        page: number = 1,
        include_adult: boolean = false,
        language: string = "en-US"
    ) {
        var options: modules.SearchTVOptions = {
            page: page,
            include_adult: include_adult,
            language: language,
        }
        let unmarshalledJson: string = await SearchTV(query,options)
        queryResult = JSON.parse(unmarshalledJson)
        queryValues = queryResult["Results"]
    }

    async function searchPeople(
        query: string,
        page: number = 1,
        include_adult: boolean = false,
        region: string = "US",
        language: string = "en-US"
    ) {
        var options: modules.SearchPeopleOptions = {
            page: page,
            include_adult: include_adult,
            region: region,
            language: language,
        }
        let unmarshalledJson: string = await SearchPeople(query,options)
        queryResult = JSON.parse(unmarshalledJson)
        queryValues = queryResult["Results"]
    }

    let searchText: string
    let queryResult: JSON
    let queryValues: JSON[] = []

    let activeSearchOptions: Map<string, Function> = new Map([
        ["Movies", searchMovies],
        ["TV Shows", searchTV],
        ["People", searchPeople]
    ])
    let activeSearchType: string = "Movies"
    let activeSearchFunction: Function = activeSearchOptions.get(activeSearchType)
    let currentSearchType: string = activeSearchType
    let currentSearchFunction: Function = activeSearchFunction
</script>

<div class="container mx-auto">
    <div class="row py-3">
        <div class="col-12">
            <div class="flex justify-center">
                <!-- Dropdown to select between movies, tv shows, and people -->
                <div class="relative inline-block text-left">
                    <div>
                        <button
                            type="button"
                            class="inline-flex justify-center w-full rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                            id="search-type-menu"
                            aria-haspopup="true"
                            aria-expanded="true"
                            on:click={() => {
                                let searchTypeItems = document.getElementById("search-type-items")
                                if (searchTypeItems.classList.contains("hidden")) {
                                    searchTypeItems.classList.remove("hidden")
                                } else {
                                    searchTypeItems.classList.add("hidden")
                                }
                            }}
                        >
                            {activeSearchType}
                            <Icon icon="material-symbols:arrow-drop-down" class="self-center"/>
                        </button>
                    </div>

                    <!--
                        Dropdown panel, show/hide based on dropdown state.

                        Entering: "transition ease-out duration-100"
                            From: "transform opacity-0 scale-95"
                            To: "transform opacity-100 scale-100"
                        Leaving: "transition ease-in duration-75"
                            From: "transform opacity-100 scale-100"
                            To: "transform opacity-0 scale-95"
                    -->
                    <div
                        class="origin-top-right absolute right-0 mt-2 w-36 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-10 hidden"
                        role="menu"
                        aria-orientation="vertical"
                        aria-labelledby="options-menu"
                        id="search-type-items"
                    >
                        <div class="py-1" role="none">
                            <button
                                on:click={() => {
                                    activeSearchType = "Movies"
                                    activeSearchFunction = activeSearchOptions.get(activeSearchType)
                                    document.getElementById("search-type-items").classList.add("hidden")
                                }}
                                class="w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                                role="menuitem"
                                >
                                Movies
                            </button>
                            <button
                            on:click={() => {
                                activeSearchType = "TV Shows"
                                activeSearchFunction = activeSearchOptions.get(activeSearchType)
                                document.getElementById("search-type-items").classList.add("hidden")
                            }}
                                class="w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                                role="menuitem"
                                >
                                TV Shows
                            </button>
                            <button
                                on:click={() => {
                                    activeSearchType = "People"
                                    activeSearchFunction = activeSearchOptions.get(activeSearchType)
                                    document.getElementById("search-type-items").classList.add("hidden")
                                }}
                                class="w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"
                                role="menuitem"
                                >
                                People
                            </button>
                        </div>
                    </div>
                </div>

                <!-- Search bar -->
                <input
                    type="text"
                    class="w-1/2 p-2 rounded-md"
                    placeholder="Search for {activeSearchType}..."
                    bind:value={searchText}
                    on:keydown={async (e) => {
                        if (e.key === "Enter") {
                            currentSearchFunction = activeSearchFunction
                            currentSearchType = activeSearchType
                            await activeSearchFunction(searchText)
                            
                        }
                    }}
                />

                <!-- Search button -->
                <button
                    class="bg-blue-500 text-white rounded-md p-2"
                    on:click={async () => {
                        currentSearchFunction = activeSearchFunction
                        currentSearchType = activeSearchType
                        await activeSearchFunction(searchText)
                    }}
                >
                    Search
                </button>
            </div>
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
                searchFunction = {currentSearchFunction}
            />
        </div>
    </div>
    {/if}

    <div class="row">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-2 px-4">
            {#each queryValues as content}
                {#if currentSearchType == "Movies"}
                    <MovieCard
                        adult = {content["adult"]}
                        backdrop_path = {content["backdrop_path"]}
                        id = {content["id"]}
                        original_title = {content["original_title"]}
                        genre_ids = {content["genre_ids"]}
                        popularity = {content["popularity"]}
                        poster_path = {content["poster_path"]}
                        release_date = {content["release_date"]}
                        title = {content["title"]}
                        overview = {content["overview"]}
                        video = {content["video"]}
                        vote_average = {content["vote_average"]}
                        vote_count = {content["vote_count"]}
                    />
                {/if}
                {#if currentSearchType == "TV Shows"}
                    <TVShowCard
                        poster_path = {content["poster_path"]}
                        popularity = {content["Popularity"]}
                        id = {content["ID"]}
                        backdrop_path = {content["backdrop_path"]}
                        vote_average = {content["vote_average"]}
                        overview = {content["overview"]}
                        first_air_date = {content["first_air_date"]}
                        origin_country = {content["origin_country"]}
                        genre_ids = {content["genre_ids"]}
                        original_language = {content["original_language"]}
                        vote_count = {content["vote_count"]}
                        name = {content["Name"]}
                        original_name = {content["original_name"]}
                    />
                {/if}
                {#if currentSearchType == "People"}
                    <PeopleCard
                        profile_path = {content["profile_path"]}
                        adult = {content["Adult"]}
                        id = {content["ID"]}
                        known_for = {content["known_for"]}
                        name = {content["Name"]}
                        popularity = {content["Popularity"]}
                    />
                {/if}
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
                searchFunction = {currentSearchFunction}
            />
        </div>
    </div>
    {/if}
</div>