import { BrowserRouter, Routes, Route, Link, useLocation, useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { CartProvider, useCart } from './context/CartContext';
import { api } from './api/client';
import type { Store } from './types';
import Home from './pages/Home';
import Category from './pages/Category';
import ProductPage from './pages/Product';
import Cart from './pages/Cart';
import './index.css';

function NavBar() {
  const { count } = useCart();
  const location = useLocation();
  const navigate = useNavigate();
  const isHome = location.pathname === '/';
  const [storeName, setStoreName] = useState('商店');

  useEffect(() => {
    api.getStore().then((s) => setStoreName(s.name)).catch(() => {});
  }, []);

  return (
    <nav className="navbar">
      <div className="nav-left">
        {!isHome && (
          <button className="nav-back" onClick={() => navigate(-1)}>←</button>
        )}
        <Link to="/" className="nav-title">{storeName}</Link>
      </div>
      <Link to="/cart" className="nav-cart">
        🛒 {count > 0 && <span className="badge">{count}</span>}
      </Link>
    </nav>
  );
}

function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/category/:id" element={<Category />} />
      <Route path="/product/:id" element={<ProductPage />} />
      <Route path="/cart" element={<Cart />} />
    </Routes>
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <CartProvider>
        <NavBar />
        <AppRoutes />
      </CartProvider>
    </BrowserRouter>
  );
}
