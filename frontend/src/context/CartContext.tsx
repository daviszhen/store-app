import { createContext, useContext, useState, useEffect } from 'react';
import type { ReactNode } from 'react';
import type { CartItem } from '../types';
import { api } from '../api/client';

interface CartContextType {
  items: CartItem[];
  loading: boolean;
  total: number;
  count: number;
  addItem: (productId: number, quantity?: number) => Promise<void>;
  updateQuantity: (id: number, quantity: number) => Promise<void>;
  removeItem: (id: number) => Promise<void>;
  refresh: () => Promise<void>;
}

const CartContext = createContext<CartContextType | null>(null);

export function CartProvider({ children }: { children: ReactNode }) {
  const [items, setItems] = useState<CartItem[]>([]);
  const [loading, setLoading] = useState(false);

  const refresh = async () => {
    setLoading(true);
    try {
      const data = await api.getCart();
      setItems(data);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refresh();
  }, []);

  const addItem = async (productId: number, quantity = 1) => {
    await api.addToCart(productId, quantity);
    await refresh();
  };

  const updateQuantity = async (id: number, quantity: number) => {
    await api.updateCartItem(id, quantity);
    await refresh();
  };

  const removeItem = async (id: number) => {
    await api.deleteCartItem(id);
    await refresh();
  };

  const total = items.reduce((sum, item) => sum + (item.product?.price || 0) * item.quantity, 0);
  const count = items.reduce((sum, item) => sum + item.quantity, 0);

  return (
    <CartContext.Provider value={{ items, loading, total, count, addItem, updateQuantity, removeItem, refresh }}>
      {children}
    </CartContext.Provider>
  );
}

export function useCart() {
  const ctx = useContext(CartContext);
  if (!ctx) throw new Error('useCart must be used within CartProvider');
  return ctx;
}
