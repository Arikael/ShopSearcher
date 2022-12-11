<!--<script lang="ts">-->
<!--    import {createShopsWithSearchTerm} from "./shop";-->

<!--    let timeout: NodeJS.Timeout-->
<!--    const delay = 300-->
<!--    const onSearchKeyUp = (e) => {-->
<!--        if (e.which === 0) {-->
<!--            return-->
<!--        }-->

<!--        shopsWithSearchTerms = []-->
<!--        if (timeout) {-->
<!--            clearTimeout(timeout)-->
<!--        }-->
<!--        timeout = setTimeout(() => {-->
<!--            shopsWithSearchTerms = createShopsWithSearchTerm(searchTerm, shops)-->
<!--        }, delay)-->
<!--    }-->
<!--</script>-->
<!--<input type="text" placeholder="enter search term" class="search-input"-->
<!--       bind:value="{searchTerm}" on:keypress="{onSearchKeyUp}"/>-->