export const state = () => ({
  counter: 1
})

export const mutations = {
  increment(state) {
    state.counter++
  },
  decrement(state) {
    state.counter--
  },
  reset(state) {
    state.counter = 0
  }
}