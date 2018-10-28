<template>
  <div>
    <button @click="$store.commit('counter/decrement')" class="waves-effect waves-light btn">-</button>
    <p>{{ counter }}</p>
    <button @click="$store.commit('counter/increment')" class="waves-effect waves-light btn">+</button>
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
