export const useSpaceStore = {
  state: {
    currentSpace: null
  },

  setCurrentSpace(space) {
    this.state.currentSpace = space
    localStorage.setItem('currentSpace', JSON.stringify(space))
  },

  getCurrentSpace() {
    if (!this.state.currentSpace) {
      const saved = localStorage.getItem('currentSpace')
      if (saved) {
        this.state.currentSpace = JSON.parse(saved)
      }
    }
    return this.state.currentSpace
  }
}
