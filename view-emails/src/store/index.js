import { createStore } from 'vuex'

export default createStore({
  state: {
    emails: [],    
    total: 0,
    from: 0,
    size:20,
    fromCount: 0,
    toCount: 0,
    text:''
  },
  getters: {
  },
  mutations: {
    setEmails(state, newValue){
      state.emails = newValue
    },
    setTotal(state, newValue){
      state.total = newValue
    },
    setSize(state, newValue){
      state.size = newValue
    },
    setText(state, newValue){
      state.text = newValue
    },
    setFrom(state, newValue){
      state.from = newValue
    },
    setFromCount(state, newValue){
      state.fromCount = newValue
    },
    setToCount(state, newValue){
      state.toCount = newValue
    },
  },
  actions: {
    async getEmails({commit, state}){
      try {        
        const response = await fetch('http://localhost:3030/search?text='+ state.text + '&from='+ state.from + '&size='+state.size)
        const data = await response.json()
        commit('setEmails', data.data)
        commit('setTotal', data.total)
        commit('setFromCount', data.fromCount)
        commit('setToCount', data.toCount)
        commit('setSize', data.size)
      } catch (error) {
        console.error(error)
      }
    },
    changeSize({ dispatch, commit }, size){
      if(size != 20 || size != 50 || size != 100 )
        commit('setSize', size)
        dispatch("getEmails")
    },
    FilterByText({dispatch, commit, state}, text){
      if(text != undefined && text != state.text )
          commit('setFrom', 0)
          commit('setText', text)
          dispatch("getEmails")
          
    },
    prevPage({dispatch, commit, state}){            
          if(state.from >= (0 + state.size)){
              commit('setFrom', state.from - state.size)      
              dispatch("getEmails")
          }
    },
    nextPage({dispatch, commit, state}){          
          if((state.from + state.size) < state.total){              
              commit('setFrom', state.from + state.size)
              dispatch("getEmails")
          }
    },
  },
  modules: {
  }
})
