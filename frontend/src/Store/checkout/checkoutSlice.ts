import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { IAddress } from "../../modules/addresses/interfaces/IAddress";
import { IProduct } from "../../modules/products/interfaces/IProduct";
import { ICarrier } from "../../modules/carriers/interfaces/ICarrier";




interface CheckoutState {
    product:IProduct
    carrier:ICarrier
    shipping: IAddress;
    paymentStatus: 'idle' | 'loading' | 'succeeded' | 'failed';
    error: string | null;
  }
  
const initialState: CheckoutState = JSON.parse(localStorage.getItem("checkoutState") ?? JSON.stringify({
  product:{ ID:0, Name: "", Description: "", Price: 0,CategoryID: 0,UserID: 0,ImageURL: ""},
  carrier:{ID:0, contact:"",name:""},
  shipping: { city: "", province: "", postalCode: "", country: "", street: "", number: 0 },
  paymentStatus: 'idle',
  error: null,
}));



const checkoutSlice = createSlice({
  name: "checkout",
  initialState,
  reducers: {
    updateShipping: (state, action: PayloadAction<IAddress>) => {
      state.shipping = action.payload;
      localStorage.setItem("checkoutState", JSON.stringify(state));
    },
    updateProduct: (state, action: PayloadAction<IProduct>) => {
      state.product = action.payload;
      localStorage.setItem("checkoutState", JSON.stringify(state));
    },
    updateCarrier: (state, action: PayloadAction<ICarrier>) => {
      state.carrier = action.payload;
      localStorage.setItem("checkoutState", JSON.stringify(state));
    },
    setPaymentStatus: (state, action: PayloadAction<'idle' | 'loading' | 'succeeded' | 'failed'>) => {
        state.paymentStatus = action.payload;
      },
    setPaymentError: (state, action: PayloadAction<string | null>) => {
        state.error = action.payload;
      },
    resetCheckout: (state) => {
      state.product={ ID:0, Name: "", Description: "", Price: 0,CategoryID: 0,UserID: 0,ImageURL: ""};
      state.carrier={ ID:0, contact:"", name:""};
      state.shipping = { city: "", province: "", postalCode: "", country: "", street: "", number: 0 };
      state.paymentStatus = 'idle';
      state.error = null;
      localStorage.removeItem("checkoutState");
    }
  }
});

export const { updateShipping,setPaymentStatus,setPaymentError, resetCheckout, updateProduct,updateCarrier} = checkoutSlice.actions;
export default checkoutSlice.reducer;