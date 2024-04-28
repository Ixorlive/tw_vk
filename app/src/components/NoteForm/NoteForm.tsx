import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { NOTE_API_CREATE, NOTE_API_EDIT } from '../../config';
import { NoteType } from '../../entity/NoteType';
import { UserType } from '../../entity/UserType';

function NoteForm({ user, note, onSave }: { user: UserType, note: NoteType | null, onSave: any }) {
    const [noteContent, setNoteContent] = useState('');
    const isNewNote = note === null;

    useEffect(() => {
        if (!isNewNote) {
            const fetchNoteDetails = async () => {
                try {
                    const response = await fetch(`${NOTE_API_EDIT}/${note.id}`);
                    if (!response.ok) {
                        throw new Error('Failed to fetch note details');
                    }
                    const data = await response.json();
                    setNoteContent(data.content);
                } catch (error) {
                    console.error('Error fetching note details:', error);
                }
            };
            fetchNoteDetails();
        }
    }, [note, isNewNote]);

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        const url = isNewNote ? NOTE_API_CREATE : `${NOTE_API_EDIT}/${note.id}`;
        const method = isNewNote ? 'POST' : 'PUT';

        try {
            const response = await fetch(url, {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ content: noteContent, user_id: user.id }),
            });
            if (!response.ok) {
                throw new Error(isNewNote ? 'Failed to create note' : 'Failed to update note');
            }
            alert(isNewNote ? 'Note created successfully' : 'Note updated successfully');
            onSave()
        } catch (error) {
            console.error(isNewNote ? 'Error creating note:' : 'Error updating note:', error);
            alert(isNewNote ? 'Failed to create note' : 'Failed to update note');
            onSave()
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Note Content:
                <textarea value={noteContent} onChange={e => setNoteContent(e.target.value)} />
            </label>
            <button type="submit">{isNewNote ? 'Create Note' : 'Update Note'}</button>
        </form>
    );
}

export default NoteForm;
