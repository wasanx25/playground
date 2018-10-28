<template>
  <div>
    <button @click="$store.commit('counter/increment')">{{ counter }}</button>
  </div>
</template>

<script>
import { mapMutations } from "vuex"

export default {
  computed: {
    counter() {
      return this.$store.state.counter.counter
    }
  }
}
</script>
