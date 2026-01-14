<template>
    <FormKit
        type="form"
        submit-label="Submit form"
        use-local-storage
        @submit="uploadImage"
    >
        <FormKit
            name="id"
            type="number"
            label="Product ID"
            help="Enter the product id to upload an image for"
            validation="required"
            :plugins="[castIdToNumber]"
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
import axios from "axios";

definePageMeta({
    title: "Admin",
});

const castIdToNumber = (node) =>
    node.hook.input((value, next) => next(Number(value)));

const uploadImage = async (fields) => {
    const data = new FormData();
    data.append("id", fields.id);
    data.append("file", fields.image[0].file);

    axios.post("http://localhost:9091/", data);
};
</script>
