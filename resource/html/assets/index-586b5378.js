import{C as W,b as v,_ as S}from"./index.js";import{P as T}from"./index-9fe0b115.js";import{d as b,a8 as c,_ as I,a9 as _,aa as e,l as o,E as a}from"./vue-5c68ae35.js";import"./antd-12f11a56.js";import"./useContentViewHeight-c2d18fda.js";import"./useWindowSizeFn-997fa1d0.js";import"./onMountedOrActivated-97a1bb6a.js";const y=b({components:{CollapseContainer:W,PageWrapper:T},setup(){const{createMessage:s,createConfirm:n,createSuccessModal:d,createInfoModal:u,createErrorModal:f,createWarningModal:m,notification:r}=v(),{info:t,success:i,warning:l,error:g}=s;function p(){s.loading("Loading...")}function C(h){n({iconType:h,title:"Tip",content:"content message..."})}function M(){d({title:"Tip",content:"content message..."})}function k(){f({title:"Tip",content:"content message..."})}function w(){m({title:"Tip",content:"content message..."})}function E(){u({title:"Tip",content:"content message..."})}function $(){r.success({message:"Tip",description:"content message..."})}return{infoMsg:t,successMsg:i,warningMsg:l,errorMsg:g,handleLoading:p,handleConfirm:C,handleSuccessModal:M,handleErrorModal:k,handleWarningModal:w,handleInfoModal:E,handleNotify:$}}});function N(s,n,d,u,f,m){const r=c("a-button"),t=c("CollapseContainer"),i=c("PageWrapper");return I(),_(i,{title:"消息示例"},{default:e(()=>[o(t,{class:"w-full h-32 bg-white rounded-md",title:"Message"},{default:e(()=>[o(r,{onClick:n[0]||(n[0]=l=>s.infoMsg("Info message")),class:"mr-2"},{default:e(()=>[a(" Info ")]),_:1}),o(r,{onClick:n[1]||(n[1]=l=>s.successMsg("Success message")),class:"mr-2",color:"success"},{default:e(()=>[a(" Success ")]),_:1}),o(r,{onClick:n[2]||(n[2]=l=>s.warningMsg("Warning message")),class:"mr-2",color:"warning"},{default:e(()=>[a(" Warning ")]),_:1}),o(r,{onClick:n[3]||(n[3]=l=>s.errorMsg("Error message")),class:"mr-2",color:"error"},{default:e(()=>[a(" Error ")]),_:1}),o(r,{onClick:s.handleLoading,class:"mr-2",type:"primary"},{default:e(()=>[a(" Loading ")]),_:1},8,["onClick"])]),_:1}),o(t,{class:"w-full h-32 mt-5 bg-white rounded-md",title:"Comfirm"},{default:e(()=>[o(r,{onClick:n[4]||(n[4]=l=>s.handleConfirm("info")),class:"mr-2"},{default:e(()=>[a(" Info ")]),_:1}),o(r,{onClick:n[5]||(n[5]=l=>s.handleConfirm("warning")),color:"warning",class:"mr-2"},{default:e(()=>[a(" Warning ")]),_:1}),o(r,{onClick:n[6]||(n[6]=l=>s.handleConfirm("success")),color:"success",class:"mr-2"},{default:e(()=>[a(" Success ")]),_:1}),o(r,{onClick:n[7]||(n[7]=l=>s.handleConfirm("error")),color:"error",class:"mr-2"},{default:e(()=>[a(" Error ")]),_:1})]),_:1}),o(t,{class:"w-full h-32 mt-5 bg-white rounded-md",title:"Modal"},{default:e(()=>[o(r,{onClick:s.handleInfoModal,class:"mr-2"},{default:e(()=>[a(" Info ")]),_:1},8,["onClick"]),o(r,{onClick:s.handleSuccessModal,color:"success",class:"mr-2"},{default:e(()=>[a(" Success ")]),_:1},8,["onClick"]),o(r,{onClick:s.handleErrorModal,color:"error",class:"mr-2"},{default:e(()=>[a(" Error ")]),_:1},8,["onClick"]),o(r,{onClick:s.handleWarningModal,color:"warning",class:"mr-2"},{default:e(()=>[a(" Warning ")]),_:1},8,["onClick"])]),_:1}),o(t,{class:"w-full h-32 mt-5 bg-white rounded-md",title:"Notification 用法与上面一致"},{default:e(()=>[o(r,{onClick:s.handleNotify,color:"success",class:"mr-2"},{default:e(()=>[a(" Success ")]),_:1},8,["onClick"])]),_:1})]),_:1})}const j=S(y,[["render",N]]);export{j as default};
