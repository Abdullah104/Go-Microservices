<template>
    <FormKit type="form" submit-label="Submit form" @submit="uploadImage">
        <FormKit
            name="id"
            type="number"
            label="Product ID"
            help="Enter the product id to upload an image for"
            :plugins="[castIdToNumber]"
            validation="required"
        />
        <FormKit
            name="image"
            type="file"
            label="File"
            help="Image to associate with the product"
            validation="required"
            accept=".png"
        />
    </FormKit>
</template>

<script setup lang="ts">
definePageMeta({
    title: "Admin",
});

const castIdToNumber = (node) =>
    node.hook.input((value, next) => next(Number(value)));

const uploadImage = async (fields) => {};
</script>
