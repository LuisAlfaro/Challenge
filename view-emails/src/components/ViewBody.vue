<template>
  <div class="border">
    <!-- header -->
    <div class="px-4 py-3 border-b border-gray-200">
      <h2 class="text-xl font-semibold text-gray-600">{{ subject }}</h2>
    </div>

    <!-- body -->
    <div class="w-full p-3">
      {{ body }}
    </div>
  </div>
</template>

<script>
export default {
  props: {
    body: {
      type: String,
      default: "",
    },
    subject: {
      type: String,
      default: "",
    },
    show: Boolean,
  },
  setup(props) {},
};
</script>
