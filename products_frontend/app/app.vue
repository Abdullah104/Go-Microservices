<template>
    <UApp>
        <h2 class="pt-10 text-center">Menu</h2>
        <UTable :data="products" :columns="columns" :loading="loading" />
    </UApp>
    <pre>{{ error }}</pre>
</template>

<script lang="ts" setup>
import type { Product } from "~/Product";
import type { TableColumn } from "#ui/components/Table.vue";

const apiLocation = "http://localhost:9090";

const columns: TableColumn<Product>[] = [
    {
        accessorKey: "name",
        header: "Name",
    },
    {
        accessorKey: "price",
        header: "Price",
    },
    {
        accessorKey: "sku",
        header: "SKU",
    },
];

const {
    data: products,
    pending: loading,
    error,
} = await useAsyncData<Product[]>("products", (_nuxtApp, { signal }) =>
    $fetch(apiLocation + "/products", { signal }),
);
</script>
