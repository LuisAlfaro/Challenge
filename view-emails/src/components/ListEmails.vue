<template>
  <div class="-mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">
    <div class="inline-block min-w-full shadow rounded-lg overflow-hidden">
      <div
        class="overflow-x-auto bg-white rounded-lg shadow overflow-y-auto relative"
        style="height: 405px"
      >
        <table
          class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
        >
          <thead>
            <tr>
              <th
                class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
              >
                Subject
              </th>
              <th
                class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
              >
                From
              </th>
              <th
                class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
              >
                To
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(email, i) in emails"
              :key="i"
              @click="getBody(i)"
              data-modal-toggle="default-modal"
            >
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                <p class="text-gray-900 whitespace-no-wrap">
                  {{ email.subject }}
                </p>
              </td>
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                <p class="text-gray-900 whitespace-no-wrap">{{ email.from }}</p>
              </td>
              <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
                <p class="text-gray-900 whitespace-no-wrap">{{ email.to }}</p>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div
        class="px-5 py-5 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between"
      >
        <span class="text-xs xs:text-sm text-gray-900">
          Showing {{ fromCount }} to {{ toCount }} of {{ total }} Entries
        </span>
        <div class="inline-flex mt-2 xs:mt-0">
          <button
            class="text-sm bg-gray-300 hover:bg-gray-400 text-gray-800 font-semibold py-2 px-4 rounded-l"
            @click="prevPage"
          >
            Prev
          </button>
          <button
            class="text-sm bg-gray-300 hover:bg-gray-400 text-gray-800 font-semibold py-2 px-4 rounded-r"
            @click="nextPage"
          >
            Next
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { onMounted, computed, ref } from "vue";
import { useStore } from "vuex";
import ViewBody from "./ViewBody.vue";
export default {
  components: { ViewBody },
  setup() {
    const store = useStore();

    const emails = computed(() => {
      return store.state.emails;
    });

    const fromCount = computed(() => {
      return store.state.fromCount;
    });

    const toCount = computed(() => {
      return store.state.toCount;
    });

    const total = computed(() => {
      return store.state.total;
    });

    let body = ref("");
    body.value = computed(() => {
      if (store.state.emails.length > 0) {
        return store.state.emails[0].body;
      }
    });

    const prevPage = () => {
      store.dispatch("prevPage");
    };

    const nextPage = () => {
      store.dispatch("nextPage");
    };

    onMounted(() => {
      store.dispatch("getEmails");
    });

    function getBody(index) {
      body.value = store.state.emails[index].body;
    }

    return {
      emails,
      body,
      fromCount,
      toCount,
      total,
      getBody,
      prevPage,
      nextPage,
    };
  },
};
</script>
