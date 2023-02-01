<template>
  <div>
    <p>
      {{ body }}
    </p>
  </div>
</template>

<script>
export default {
  props: {
    body: {
      type: String,
      default: "",
    },
  },
  setup(props) {},
};
</script>
