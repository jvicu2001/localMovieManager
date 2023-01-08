<script lang="ts">
    export let currentPage: number
    export let totalPages: number
    export let searchText: string
    export let searchFunction: Function
    export let sidePages: number = 5
    
    // Icons
    import Icon from '@iconify/svelte';
</script>

<div class="flex justify-center">
    {#if currentPage > 1}
        <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, 1)}>
            <Icon icon="material-symbols:first-page" class="mx-auto"/>
        </button>
        <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, currentPage - 1)}>
            <Icon icon="material-symbols:chevron-left" class="mx-auto"/>
        </button>
    {/if}

    <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, 1)}>1</button>
    
    {#if totalPages > 1}
        {#if (currentPage - sidePages) > 1}
            
            <div class="text-white ring-2 rounded-md w-10 p-2">...</div>
        {/if}
        {#each Array(totalPages) as page, i}
            {#if (i + 1) > (currentPage - sidePages) && (i + 1) < (currentPage + sidePages) && ((i+1) !== 1) && ((i+1) !== totalPages)}
                <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, i+1)}>{i + 1}</button>
            {/if}
        {/each}
        {#if (currentPage + sidePages) < totalPages}
            <div class="text-white ring-2 rounded-md w-10 p-2">...</div>
            
        {/if} 

        <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, totalPages)}>{totalPages}</button>
    {/if}

    
        
    {#if currentPage < totalPages}
        <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, currentPage + 1)}>
            <Icon icon="material-symbols:chevron-right" class="mx-auto"/>
        </button>
        <button class="text-white ring-2 rounded-md w-10 p-2" on:click={() => searchFunction(searchText, totalPages)}>
            <Icon icon="material-symbols:last-page" class="mx-auto"/>
        </button>
    {/if}
    
</div>

<style>

</style>