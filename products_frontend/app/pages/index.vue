<template>
    <div>
        <UTable :data="products" :columns="columns" :loading="loading" />
        <pre>{{ error }}</pre>
    </div>
</template>

<script setup lang="ts">
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

definePageMeta({
    title: "Main",
});
</script>
