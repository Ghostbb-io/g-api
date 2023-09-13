var Q=Object.defineProperty,X=Object.defineProperties;var Y=Object.getOwnPropertyDescriptors;var B=Object.getOwnPropertySymbols;var Z=Object.prototype.hasOwnProperty,ee=Object.prototype.propertyIsEnumerable;var $=(e,u,m)=>u in e?Q(e,u,{enumerable:!0,configurable:!0,writable:!0,value:m}):e[u]=m,f=(e,u)=>{for(var m in u||(u={}))Z.call(u,m)&&$(e,m,u[m]);if(B)for(var m of B(u))ee.call(u,m)&&$(e,m,u[m]);return e},H=(e,u)=>X(e,Y(u));import{c as F,u as n,d as te,G as se,n as oe,l as g,i as ne,z as le,ad as J,E as ae}from"./vue-5c68ae35.js";import{c as z}from"./componentMap-c5f962c4.js";import{s as re,m as x,O as N,k as ce,x as E,a1 as ie,a2 as de}from"./index.js";import{N as fe,c as K,s as ue}from"./helper-736f1539.js";import{O as G,a1 as me,$ as he,d as pe,k as be}from"./antd-12f11a56.js";function ge(e,u){return F(()=>{const m=n(e),{labelCol:j={},wrapperCol:L={}}=m.itemProps||{},{labelWidth:O,disabledLabelWidth:p}=m,{labelWidth:w,labelCol:I,wrapperCol:k,layout:q}=n(u);if(!w&&!O&&!I||p)return j.style={textAlign:"left"},{labelCol:j,wrapperCol:L};let P=O||w;const D=f(f({},I),j),_=f(f({},k),L);return P&&(P=re(P)?`${P}px`:P),{labelCol:f({style:{width:P}},D),wrapperCol:f({style:{width:q==="vertical"?"100%":`calc(100% - ${P})`}},_)}})}function W(e){return typeof e=="function"||Object.prototype.toString.call(e)==="[object Object]"&&!ne(e)}const Ie=te({name:"BasicFormItem",inheritAttrs:!1,props:{schema:{type:Object,default:()=>({})},formProps:{type:Object,default:()=>({})},allDefaultValues:{type:Object,default:()=>({})},formModel:{type:Object,default:()=>({})},setFormModel:{type:Function,default:null},tableAction:{type:Object},formActionType:{type:Object},isAdvanced:{type:Boolean}},setup(e,{slots:u}){const{t:m}=ce(),{schema:j,formProps:L}=se(e),O=ge(j,L),p=F(()=>{const{allDefaultValues:s,formModel:t,schema:a}=e,{mergeDynamicData:r}=e.formProps;return{field:a.field,model:t,values:f(f(f({},r),s),t),schema:a}}),w=F(()=>{var l;const{schema:s,tableAction:t,formModel:a,formActionType:r}=e;let{componentProps:o={}}=s;return x(o)&&(o=(l=o({schema:s,tableAction:t,formModel:a,formActionType:r}))!=null?l:{}),s.component==="Divider"&&(o=Object.assign({type:"horizontal"},{orientation:"left",plain:!0},o)),o}),I=F(()=>{const{disabled:s}=e.formProps,{dynamicDisabled:t}=e.schema,{disabled:a=!1}=n(w);let r=!!s||a;return N(t)&&(r=t),x(t)&&(r=t(n(p))),r});function k(){const{show:s,ifShow:t}=e.schema,{showAdvancedButton:a}=e.formProps,r=a&&N(e.isAdvanced)?e.isAdvanced:!0;let o=!0,l=!0;return N(s)&&(o=s),N(t)&&(l=t),x(s)&&(o=s(n(p))),x(t)&&(l=t(n(p))),o=o&&r,{isShow:o,isIfShow:l}}function q(){var v;const{rules:s=[],component:t,rulesMessageJoinLabel:a,label:r,dynamicRules:o,required:l}=e.schema;if(x(o))return o(n(p));let c=pe(s);const{rulesMessageJoinLabel:S}=e.formProps,A=Reflect.has(e.schema,"rulesMessageJoinLabel")?a:S,y=K(t)+`${A?r:""}`;function b(d,i){const C=d.message||y;return i===void 0||de(i)||Array.isArray(i)&&i.length===0||typeof i=="string"&&i.trim()===""||typeof i=="object"&&Reflect.has(i,"checked")&&Reflect.has(i,"halfChecked")&&Array.isArray(i.checked)&&Array.isArray(i.halfChecked)&&i.checked.length===0&&i.halfChecked.length===0?Promise.reject(C):Promise.resolve()}const h=x(l)?l(n(p)):l;h&&(!c||c.length===0?c=[{required:h,validator:b}]:c.findIndex(i=>Reflect.has(i,"required"))===-1&&c.push({required:h,validator:b}));const R=c.findIndex(d=>Reflect.has(d,"required")&&!Reflect.has(d,"validator"));if(R!==-1){const d=c[R],{isShow:i}=k();if(i||(d.required=!1),t){Reflect.has(d,"type")||(d.type=t==="InputNumber"?"number":"string"),d.message=d.message||y,(t.includes("Input")||t.includes("Textarea"))&&(d.whitespace=!0);const C=(v=n(w))==null?void 0:v.valueFormat;ue(d,t,C)}}const M=c.findIndex(d=>d.max);return M!==-1&&!c[M].validator&&(c[M].message=c[M].message||m("component.form.maxTip",[c[M].max])),c}function P(){var i;const{renderComponentContent:s,component:t,field:a,changeEvent:r="change",valueField:o}=e.schema,l=t&&["Switch","Checkbox"].includes(t),c=`on${be(r)}`,S={[c]:(...C)=>{const[T]=C;h[c]&&h[c](...C);const V=T?T.target:null,U=V?l?V.checked:V.value:T;e.setFormModel(a,U,e.schema)}},A=z.get(t),{autoSetPlaceHolder:y,size:b}=e.formProps,h=H(f({allowClear:!0,getPopupContainer:C=>C.parentNode,size:b},n(w)),{disabled:n(I)});!h.disabled&&y&&t!=="RangePicker"&&t&&(h.placeholder=((i=n(w))==null?void 0:i.placeholder)||K(t)),h.codeField=a,h.formValues=n(p);const M={[o||(l?"checked":"value")]:e.formModel[a]},v=f(f(f({},h),S),M);if(!s)return g(A,v,null);const d=x(s)?f({},s(n(p),{disabled:n(I)})):{default:()=>s};return g(A,v,W(d)?d:{default:()=>[d]})}function D(){const{label:s,helpMessage:t,helpComponentProps:a,subLabel:r}=e.schema,o=r?g("span",null,[s,ae(" "),g("span",{class:"text-secondary"},[r])]):s,l=x(t)?t(n(p)):t;return!l||Array.isArray(l)&&l.length===0?o:g("span",null,[o,g(ie,J({placement:"top",class:"mx-1",text:l},a),null)])}function _(){const{itemProps:s,slot:t,render:a,field:r,suffix:o,component:l}=e.schema,{labelCol:c,wrapperCol:S}=n(O),{colon:A}=e.formProps,y={disabled:n(I)};if(l==="Divider"){let b;return g(G,{span:24},{default:()=>[g(me,n(w),W(b=D())?b:{default:()=>[b]})]})}else{const b=()=>t?E(u,t,n(p),y):a?a(n(p),y):P(),h=!!o,R=x(o)?o(n(p)):o;return fe.includes(l)&&e.schema&&(e.schema.itemProps=f({autoLink:!1},e.schema.itemProps)),g(he.Item,J({name:r,colon:A,class:{"suffix-item":h}},s,{label:D(),rules:q(),labelCol:c,wrapperCol:S}),{default:()=>[g("div",{style:"display:flex"},[g("div",{style:"flex:1;"},[b()]),h&&g("span",{class:"suffix"},[R])])]})}}return()=>{let s;const{colProps:t={},colSlot:a,renderColContent:r,component:o}=e.schema;if(!z.has(o))return null;const{baseColProps:l={}}=e.formProps,c=f(f({},l),t),{isIfShow:S,isShow:A}=k(),y=n(p),b={disabled:n(I)};return S&&oe(g(G,c,W(s=(()=>a?E(u,a,y,b):r?r(y,b):_())())?s:{default:()=>[s]}),[[le,A]])}}});export{Ie as _};