export const state = () => ({
  counter: 0
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
  },
  plusTen(state) {
    state.counter += 10
  },
  minusTen(state) {
    state.counter -= 10
  }
}