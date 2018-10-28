<template>
  <div id="counter_content">
    <button @click="$store.commit('counter/minusTen')" class="waves-effect waves-light btn">-10</button>
    <button @click="$store.commit('counter/decrement')" class="waves-effect waves-light btn">-</button>
    <p>{{ counter }}</p>
    <button @click="$store.commit('counter/increment')" class="waves-effect waves-light btn">+</button>
    <button @click="$store.commit('counter/plusTen')" class="waves-effect waves-light btn">+10</button>
    <button id="reset" @click="$store.commit('counter/reset')" class="waves-effect waves-light btn-small blue-grey lighten-3">Reset</button>
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

<style>
#counter_content {
  text-align: center;
}

button {
  display: inline-block;
}

#counter_content p {
  font-size: 20px;
  display: inline-block;
  margin: 0 20px;
}

#reset {
  display: block;
  margin: 1rem auto;
  font-size: 15px;
}
</style>
