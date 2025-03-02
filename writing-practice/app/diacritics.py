import streamlit as st
from PIL import Image

st.set_page_config(
    page_title="Greek Diacritic Typing Practice",
    page_icon="⌨️",
    layout="centered",
)

st.title("⌨️ Welcome to Greek Diacritic Typing Practice!")
st.subheader("Use this page to practice typing Greek diacritic marks using the Greek Keyboard layout.")
st.divider()

# Display the Greek Keyboard Layout
st.subheader("Greek Keyboard Layout with Diacritic Marks")
st.image("./img/polytonic_keyboard.png", caption="Greek Keyboard Layout", use_container_width=True)

# Instructions
st.markdown("""
**How to Practice**:
1. Use the Greek Keyboard layout above.
2. Try typing the given diacritic mark by selecting from the options below.
3. Type the correct letter using your keyboard.
4. After typing, press Enter to check if you typed the correct diacritic mark.
""")

# Diacritic marks organized by type
diacritic_marks = [
    # Basic accents (acute, grave, circumflex)
    'ά', 'ὰ', 'ᾶ',     # alpha
    'έ', 'ὲ',          # epsilon
    'ή', 'ὴ', 'ῆ',     # eta
    'ί', 'ὶ', 'ῖ',     # iota
    'ύ', 'ὺ', 'ῦ',     # upsilon
    'ό', 'ὸ',          # omicron
    'ώ', 'ὼ', 'ῶ',     # omega

    # Breathing marks (smooth, rough)
    'ἀ', 'ἁ',          # alpha
    'ἐ', 'ἑ',          # epsilon
    'ἠ', 'ἡ',          # eta
    'ἰ', 'ἱ',          # iota
    'ὐ', 'ὑ',          # upsilon
    'ὀ', 'ὁ',          # omicron
    'ὠ', 'ὡ',          # omega

    # Breathing + accent combinations
    'ἄ', 'ἅ', 'ἂ', 'ἃ', 'ἆ', 'ἇ',    # alpha
    'ἔ', 'ἕ', 'ἒ', 'ἓ',               # epsilon
    'ἤ', 'ἥ', 'ἢ', 'ἣ', 'ἦ', 'ἧ',    # eta
    'ἴ', 'ἵ', 'ἲ', 'ἳ', 'ἶ', 'ἷ',    # iota
    'ὔ', 'ὕ', 'ὒ', 'ὓ', 'ὖ', 'ὗ',    # upsilon
    'ὄ', 'ὅ', 'ὂ', 'ὃ',               # omicron
    'ὤ', 'ὥ', 'ὢ', 'ὣ', 'ὦ', 'ὧ',    # omega

    # Iota subscript combinations (only for α, η, ω)
    'ᾳ', 'ῃ', 'ῳ',                    # basic iota subscript
    'ᾴ', 'ᾲ', 'ᾷ',                    # alpha + iota subscript + accents
    'ῄ', 'ῂ', 'ῇ',                    # eta + iota subscript + accents
    'ῴ', 'ῲ', 'ῷ',                    # omega + iota subscript + accents

    # Iota subscript + breathing combinations
    'ᾀ', 'ᾁ', 'ᾄ', 'ᾅ', 'ᾂ', 'ᾃ', 'ᾆ', 'ᾇ',    # alpha
    'ᾐ', 'ᾑ', 'ᾔ', 'ᾕ', 'ᾒ', 'ᾓ', 'ᾖ', 'ᾗ',    # eta
    'ᾠ', 'ᾡ', 'ᾤ', 'ᾥ', 'ᾢ', 'ᾣ', 'ᾦ', 'ᾧ'     # omega
]

# Choose a diacritic mark to practice
diacritic_mark_to_practice = st.selectbox(
    "Select a diacritic mark to practice:",
    diacritic_marks
)

# Input box for typing the selected diacritic mark
user_input = st.text_input(f"Type the diacritic mark: {diacritic_mark_to_practice}")

# Check if the input is correct
if user_input:
    if user_input == diacritic_mark_to_practice:
        st.success("Correct! 🎉 Keep practicing.")
    else:
        st.error("Oops! Try again. 😅")
        
# Footer with reminders
st.divider()
st.markdown("""
🎯 **Keep Practicing!**
- Switch between different diacritic marks to sharpen your typing skills.
- Practice regularly and track your progress!

""")
