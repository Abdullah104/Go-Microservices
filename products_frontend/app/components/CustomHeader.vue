<template>
    <UHeader :toggle="false">
        <template #title>Coffee Shop</template>
        <UNavigationMenu :items="items" />
    </UHeader>
</template>

<script setup lang="ts">
import type { NavigationMenuItem } from "@nuxt/ui";

const items = computed<NavigationMenuItem[]>(() => [
    {
        label: "Home",
        active: true,
    },
    {
        label: "Admin",
    },
]);
</script>
