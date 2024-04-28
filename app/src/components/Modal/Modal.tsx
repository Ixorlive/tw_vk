import React from 'react';
import ReactDOM from 'react-dom';

function Modal({ children, onClose }: { children: React.ReactNode; onClose: () => void }) {
    return ReactDOM.createPortal(
        <div className="modal-overlay" onClick={onClose}
            style={{
                position: 'fixed', top: 0, left: 0, right: 0, bottom: 0,
                backgroundColor: 'rgba(0,0,0,0.5)', display: 'flex', alignItems: 'center', justifyContent: 'center'
            }}>
            <div className="modal-content" onClick={e => e.stopPropagation()}
                style={{
                    backgroundColor: 'white', padding: '20px', borderRadius: '5px', maxWidth: '500px', width: '100%'
                }}>
                <button onClick={onClose} style={{ position: 'absolute', top: '10px', right: '10px' }}>Close</button>
                {children}
            </div>
        </div>,
        document.body
    );
}

export default Modal;
