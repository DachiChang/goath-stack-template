const themeData = () => {
  const defaultTheme = 'light';
  const switchedTheme = 'dark';
  const saved = localStorage.getItem('theme') || defaultTheme;
  return {
    default: defaultTheme,
    switched: switchedTheme,
    theme: saved,
    get themeToggle() {
      return this.theme === this.switched;
    },
    set themeToggle(val) {
      this.theme = val ? this.switched : this.default;
      localStorage.setItem('theme', this.theme);
    },
  };
};

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
