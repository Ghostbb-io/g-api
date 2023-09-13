var V=Object.defineProperty,k=Object.defineProperties;var z=Object.getOwnPropertyDescriptors;var v=Object.getOwnPropertySymbols;var G=Object.prototype.hasOwnProperty,L=Object.prototype.propertyIsEnumerable;var P=(t,e,r)=>e in t?V(t,e,{enumerable:!0,configurable:!0,writable:!0,value:r}):t[e]=r,o=(t,e)=>{for(var r in e||(e={}))G.call(e,r)&&P(t,r,e[r]);if(v)for(var r of v(e))L.call(e,r)&&P(t,r,e[r]);return t},C=(t,e)=>k(t,z(e));import{f as q,v as F,x as H,C as J,m as O,w as K}from"./index.js";import{a3 as w,a4 as Q}from"./antd-12f11a56.js";import{d as T,f as U,c as d,u as s,l as u,ad as X,i as Y,G as Z}from"./vue-5c68ae35.js";function tt(t){return typeof t=="function"||Object.prototype.toString.call(t)==="[object Object]"&&!Y(t)}const et={useCollapse:{type:Boolean,default:!0},title:{type:String,default:""},size:{type:String,validator:t=>["small","default","middle",void 0].includes(t),default:"small"},bordered:{type:Boolean,default:!0},column:{type:[Number,Object],default:()=>({xxl:4,xl:3,lg:3,md:3,sm:2,xs:1})},collapseOptions:{type:Object,default:null},schema:{type:Array,default:()=>[]},data:{type:Object}},nt=T({name:"Description",props:et,emits:["register"],setup(t,{slots:e,emit:r}){const m=U(null),{prefixCls:S}=q("description"),_=F(),g=d(()=>o(o({},t),s(m))),f=d(()=>C(o({},s(g)),{title:void 0})),E=d(()=>!!s(g).title),I=d(()=>o({canExpand:!1},s(f).collapseOptions)),M=d(()=>o(o({},s(_)),s(f)));function N(n){m.value=o(o({},s(m)),n)}function W({label:n,labelMinWidth:c,labelStyle:a}){if(!a&&!c)return n;const l=C(o({},a),{minWidth:`${c}px `});return u("div",{style:l},[n])}function $(){const{schema:n,data:c}=s(f);return s(n).map(a=>{const{render:l,field:h,span:B,show:x,contentMinWidth:b}=a;if(x&&O(x)&&!x(c))return null;const D=()=>{var j;const i=(j=s(f))==null?void 0:j.data;if(!i)return null;const p=Q(i,h);return p&&!Z(i).hasOwnProperty(h)?O(l)?l("",i):"":O(l)?l(p,i):p!=null?p:""},R=b;return u(w.Item,{label:W(a),key:h,span:B},{default:()=>{if(!b)return D();const i={minWidth:`${R}px`};return u("div",{style:i},[D()])}})}).filter(a=>!!a)}const y=()=>{let n;return u(w,X({class:`${S}`},s(M)),tt(n=$())?n:{default:()=>[n]})},A=()=>{const n=t.useCollapse?y():u("div",null,[y()]);if(!t.useCollapse)return n;const{canExpand:c,helpMessage:a}=s(I),{title:l}=s(g);return u(J,{title:l,canExpan:c,helpMessage:a},{default:()=>n,action:()=>H(e,"action")})};return r("register",{setDescProps:N}),()=>s(E)?A():y()}}),it=K(nt);export{it as D};
