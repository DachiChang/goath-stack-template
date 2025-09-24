function demoConstructing() {
  return {
    fullText: 'Constructing ...',
    showText: '',
    i: 0,
    init() {
      this.typingEffect();
    },
    typingEffect() {
      if (this.i < this.fullText.length) {
        this.showText += this.fullText[this.i];
        this.i++;
        setTimeout(() => {
          this.typingEffect();
        }, 100);
      }
    },
  };
}
