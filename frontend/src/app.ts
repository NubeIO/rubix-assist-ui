import moment from "moment";


// import { Person } from "./wailsjs/go/models";

function getTime(): string {
  const now = moment();
  now.format("hh:mm:ss K"); // 1:00:00 PM
  return now.format("HH:mm:ss");
}

function greet(n: string) {

  // window.runtime.EventsEmit()


  // window.runtime.EventsEmit("terminal-echo", getTime())
  // const p = new Person();
  // p.name =n
  // console.log(p)
  // p.age = 42;
  // let t = ""
}

export default {
  timestamp: "",
  props: {
    source: String,
  },
  mounted() {
    // window.runtime.EventsOn("os-time", (val) => {
    //   console.log(val, "time");
    //   this.timestamp = val;
    // });
  },
  data() {
    return {
      drawer: null,
      timestamp: this.timestamp,
    };
  },
  created() {
    setInterval(() => {
      console.log("created");
      greet(" is cool");
    }, 15000);
  },
  methods: {},
};
