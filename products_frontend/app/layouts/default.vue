<template>
    <div>
        <h1 class="pt-10 text-center">{{ route.meta.title }}</h1>
        <slot />
    </div>
</template>

<script lang="ts" setup>
const route = useRoute();
</script>
