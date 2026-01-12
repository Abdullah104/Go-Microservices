<template>
    <UHeader :toggle="false">
        <template #title>Coffee Shop</template>
        <UNavigationMenu :items="items" />
    </UHeader>
</template>

<script setup lang="ts">
import type { NavigationMenuItem } from "@nuxt/ui";

const route = useRoute();

const items = computed<NavigationMenuItem[]>(() => [
    {
        label: "Home",
        active: route.name === "index",
        to: "/",
    },
    {
        label: "Admin",
        to: "/admin",
        active: route.name === "admin",
    },
]);
</script>
