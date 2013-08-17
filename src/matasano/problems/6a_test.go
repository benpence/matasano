package problems

import (
	"testing"

    "matasano/repeatingxor"
)

func Test6a(t *testing.T) {
    plaintext := []byte(
        "Eighteen-year-old Kaitlyn Hunt, charged with a crime for having sex with a 14-year-old girl, rejected a deal Friday that would have required her to plead" +
        "guilty to child abuse, according to Hunt's attorney.  Hunt was charged with two felony counts of lewd and lascivious battery after the parents of the" +
        "14-year-old went to authorities. Hunt's family says their relationship was consensual, but in Florida a person under the age of 16 is not legally able to" +
        "give consent to sex.  If Hunt is convicted, she could go to prison for 15 years -- a reality that touched off a maelstrom of controversy across the" +
        "country this past week. The case became widely known when Hunt's family began an online campaign in defense of their daughter.  The plea deal from the" +
        "Indian River County prosecutor's office would have required Hunt to plead guilty to felony child abuse, spend two years 'on community control,' which" +
        "usually involves strict supervision, followed by one year of probation. According to the plea deal document, during her probation, Hunt would have had" +
        "to agree to stay away from the 14-year-old, and to provide her probation officer with immediate access to her Internet and telephone communication." )

    key := []byte{ 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 2, 1 }

    if ciphertext, err := repeatingxor.Crypt(plaintext, key); err != nil {
        t.Error("Encrypting", plaintext, "failed")

    } else if computedKey, err := repeatingxor.FindKey(ciphertext); err != nil {
        t.Error("repeatingxor.FindKey returned error:", err)

    } else if result, err := repeatingxor.Crypt(ciphertext, computedKey); err != nil {
        t.Error("repeatingxor.Crypt returned error:", err)

    } else if string(result) != string(plaintext) {
        t.Error(string(result), "!=", string(plaintext))
	}
}
