<template>
  <div class="my-3 flex sm:flex-row flex-col">
    <div class="flex flex-row mb-1 sm:mb-0">
      <div class="relative">
        <select
          class="h-full rounded-l border block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
          v-model="size"
          @change="changeSize()"
        >
          <option value="20">20</option>
          <option value="50">50</option>
          <option value="100">100</option>
        </select>
        <div
          class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
        >
          <svg
            class="fill-current h-4 w-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          >
            <path
              d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
            />
          </svg>
        </div>
      </div>
    </div>
    <div class="block relative">
      <span class="h-full absolute inset-y-0 left-0 flex items-center pl-2">
        <svg viewBox="0 0 24 24" class="h-4 w-4 fill-current text-gray-500">
          <path
            d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z"
          ></path>
        </svg>
      </span>
      <input
        type="text"
        name=""
        id=""
        placeholder="Search..."
        class="appearance-none rounded-r rounded-l sm:rounded-l-none border border-gray-400 border-b block pl-8 pr-6 py-2 w-full bg-white text-sm placeholder-gray-400 text-gray-700 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700 focus:outline-none"
        v-model="text"
      />
    </div>
    <div class="block relative">
      <input
        type="button"
        value="Buscar"
        class="bg-blue-500 p-2 rounded-tr-lg rounded-br-lg text-white font-semibold hover:bg-blue-800 transition-colors"
        @click="filterByText()"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { useStore } from "vuex";
const store = useStore();
const text = ref("");

const selected = ref("A");
let size = ref(20);

size.value = store.state.size;

const filterByText = () => {
  store.dispatch("FilterByText", text.value);
};
const changeSize = () => {
  store.dispatch("changeSize", size.value);
};
</script>

<style></style>
