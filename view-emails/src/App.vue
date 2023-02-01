<template>
  <div class="container mx-auto px-4 sm:px-8">
    <div class="py-8">
      <div>
        <h1 class="mb-4 text-3xl font-bold text-blue-600 underline">
          View Emails
        </h1>
      </div>
      <FilterByText />
      <ListEmails />
    </div>
  </div>
</template>

<script>
import FilterByText from "./components/FilterByText.vue";
import ListEmails from "./components/ListEmails.vue";
export default {
  name: "App",
  components: { FilterByText, ListEmails },
};
</script>

<style></style>
