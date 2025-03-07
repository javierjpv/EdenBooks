import { Route, Routes } from "react-router";
import { MyProductPage } from "./modules/products/pages/MyProductPage";
import { ProductDetailPage } from "./modules/products/pages/ProductDetailPage";
import { NewProductPage } from "./modules/products/pages/NewProductPage";
import { EditProductPage } from "./modules/products/pages/EditProductPage";
import { Auth } from "./modules/users/pages/Auth";
import { ProtectedRoute } from "./routes/ProtectedRoute";
import { ShippingPage } from "./modules/checkout/pages/ShippingPage";
import { PaymentPage } from "./modules/checkout/pages/PaymentPage";
import { SuccessPage } from "./modules/checkout/pages/SuccessPage";
import { CancelPage } from "./modules/checkout/pages/CancelPage";
import { NotificationPage } from "./modules/notifications/pages/NotificationPage";
import { OrderPage } from "./modules/orders/pages/OrderPage";
import { ChatPage } from "./modules/chats/pages/ChatPage";
import { ChatDetail } from "./modules/chats/components/ChatDetail";
import { FavoritesPage } from "./modules/products/pages/FavoritesPage";


export const App = () => {
  return (
    <Routes>
      <Route path="/auth" element={<Auth />} />
      <Route element={<ProtectedRoute />}>
        <Route path="/" element={<MyProductPage />} />
        <Route path="/products" element={<MyProductPage />} />
        <Route path="/products/new" element={<NewProductPage />} />
        <Route path="/products/:id" element={<ProductDetailPage />} />
        <Route path="/products/edit/:id" element={<EditProductPage />} />

        <Route path="/notifications" element={<NotificationPage />} />

        <Route path="/checkout/shipping" element={<ShippingPage />} />
        <Route path="/checkout/payment" element={<PaymentPage />} />
        <Route path="/checkout/success" element={<SuccessPage />} />
        <Route path="/checkout/cancel" element={<CancelPage />} />

        <Route path="/orders" element={<OrderPage />} />

        <Route path="/chats" element={<ChatPage />} />
        <Route path="/chats/:id" element={<ChatDetail />} />

        <Route path="/favorites" element={<FavoritesPage />} />
      </Route>
    </Routes>
  );
};
