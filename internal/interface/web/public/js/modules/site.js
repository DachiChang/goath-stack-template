const themeData = () => ({
  default: 'light',
  switched: 'dark',
  theme: localStorage.getItem('theme') || this.default,
  get themeToggle() {
    return this.theme === this.switched;
  },
  set themeToggle(val) {
    this.theme = val ? this.switched : this.default;
    localStorage.setItem('theme', this.theme);
  },
});

const demoConstructing = () => ({
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
});

export default {
  themeData,
  demoConstructing,
};
