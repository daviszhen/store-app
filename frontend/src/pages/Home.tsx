import { useEffect, useState, useRef, useCallback, useMemo } from 'react';
import { Link } from 'react-router-dom';
import type { Store, Category, Product, CartItem } from '../types';
import { api } from '../api/client';
import { useCart } from '../context/CartContext';

export default function Home() {
  const [store, setStore] = useState<Store | null>(null);
  const [categories, setCategories] = useState<Category[]>([]);
  const [products, setProducts] = useState<Product[]>([]);
  const { items: cart, addItem, updateQuantity, removeItem } = useCart();
  const [activeCat, setActiveCat] = useState(0);
  const contentRef = useRef<HTMLDivElement>(null);
  const sectionRefs = useRef<Map<number, HTMLDivElement>>(new Map());

  useEffect(() => {
    api.getStore().then(setStore);
    api.getCategories().then(setCategories);
    api.getProducts().then(setProducts);
  }, []);

  const cartMap = useMemo(() => {
    const map: Record<number, CartItem> = {};
    cart.forEach((item) => { map[item.product_id] = item; });
    return map;
  }, [cart]);

  const handleAdd = useCallback((e: React.MouseEvent, productId: number) => {
    e.preventDefault();
    e.stopPropagation();
    addItem(productId, 1);
  }, [addItem]);

  const handleIncrease = useCallback((e: React.MouseEvent, cartItemId: number, qty: number) => {
    e.preventDefault();
    e.stopPropagation();
    updateQuantity(cartItemId, qty + 1);
  }, [updateQuantity]);

  const handleDecrease = useCallback((e: React.MouseEvent, cartItemId: number, qty: number) => {
    e.preventDefault();
    e.stopPropagation();
    if (qty <= 1) {
      removeItem(cartItemId);
    } else {
      updateQuantity(cartItemId, qty - 1);
    }
  }, [removeItem, updateQuantity]);

  const scrollToCategory = (catId: number) => {
    setActiveCat(catId);
    const el = sectionRefs.current.get(catId);
    if (el) {
      el.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }
  };

  // Group products by category
  const grouped = categories.map((cat) => ({
    ...cat,
    items: products.filter((p) => p.category_id === cat.id),
  })).filter((g) => g.items.length > 0);

  // Detect which category is in view on scroll
  const handleScroll = () => {
    const container = contentRef.current;
    if (!container) return;
    const scrollTop = container.scrollTop + 10;
    for (const g of grouped) {
      const el = sectionRefs.current.get(g.id);
      if (el && el.offsetTop <= scrollTop) {
        setActiveCat(g.id);
      }
    }
  };

  if (!store) return <div className="loading">加载中...</div>;

  return (
    <div className="page">
      {/* Header */}
      <div className="header" style={{ background: store.theme }}>
        <h1>{store.name}</h1>
        {store.notice && <p className="notice">{store.notice}</p>}
      </div>

      {/* Left categories + Right products */}
      <div className="shop-layout">
        <div className="shop-sidebar">
          {grouped.map((cat) => (
            <div
              key={cat.id}
              className={`sidebar-item ${activeCat === cat.id ? 'active' : ''}`}
              onClick={() => scrollToCategory(cat.id)}
            >
              <span className="sidebar-icon">{cat.icon}</span>
              <span>{cat.name}</span>
            </div>
          ))}
        </div>

        <div className="shop-content" ref={contentRef} onScroll={handleScroll}>
          {grouped.map((cat) => (
            <div
              key={cat.id}
              ref={(el) => { if (el) sectionRefs.current.set(cat.id, el); }}
              className="cat-section"
            >
              <div className="cat-section-title">{cat.icon} {cat.name}</div>
              <div className="product-grid">
                {cat.items.map((p) => {
                  const ci = cartMap[p.id];
                  return (
                    <div key={p.id} className="product-card">
                      <Link to={`/product/${p.id}`} className="product-card-link">
                        <div className="product-image">📦</div>
                        <div className="product-info">
                          <div className="product-name">{p.name}</div>
                          <div className="product-price">¥{p.price}</div>
                        </div>
                      </Link>
                      {ci ? (
                        <div className="cart-controls" onClick={(e) => e.preventDefault()}>
                          <button
                            className="qty-btn"
                            onClick={(e) => handleDecrease(e, ci.id, ci.quantity)}
                          >−</button>
                          <span className="qty-num">{ci.quantity}</span>
                          <button
                            className="qty-btn"
                            onClick={(e) => handleIncrease(e, ci.id, ci.quantity)}
                          >＋</button>
                        </div>
                      ) : (
                        <button
                          className="add-cart-single"
                          onClick={(e) => handleAdd(e, p.id)}
                        >＋</button>
                      )}
                    </div>
                  );
                })}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
